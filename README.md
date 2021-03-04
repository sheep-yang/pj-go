•Kubernetes的声明式API对象和控制器模型

•Kubernetes API编程范式

•如何编写一个自定义控制器

声明式API对象的编程范式

 API对象的组织方式
 
API对象在etcd里的完整资源路径是由 Group（API组）、Version（API版本）和Resource（API资源类型）三部分组成。
Kubernetes创建资源对象的流程：
首先，Kubernetes读取用户提交的yaml文件
然后，Kubernetes去匹配yaml文件中API对象的组
再次，Kubernetes去匹配yaml文件中API对象的版本号
最后，Kubernetes去匹配yaml文件中API对象的资源类型
因此，我们需要根据需求，先进行自定义资源(CRD - Custom Resource Definition)，它将包括API对象组、版本号、资源类型：


apiVersion: apiextensions.k8s.io/v1beta1

kind: CustomResourceDefinition

metadata:
    name: myresources.network
    
spec: 
    group: network
    
    version: v1
    names:
      kind: Vpc
      plural: myresources
    scope: Namespaced  
    shortNames:
    - vpc
 
 - 
自定义资源控制器
控制器与APIServer通信
Informer是APIServer与Kubernetes互相通信的桥梁，它通过Reflector实现ListAndWatch方法来“获取”和“监听”对象实例的变化。
每当APIServer接收到创建、更新和删除实例的请求，Refector都会收到“事件通知”，然后将变更的事件推送到先进先出的队列中。
Informer会不断从上一队列中读取增量，然后根据增量事件的类型创建或者更新本地对象的缓存。Informer会根据事件类型触发事先定义好的ResourceEventHandler（具体为AddFunc、UpdatedFunc和DeleteFunc，分别对应API对象的“添加”、“更新”和“删除”事件），同时每隔一定的时间，Informer也会对本地的缓存进行一次强制更新。
WorkQueue同步Informer和控制循环(Control Loop)交互的数据
Controller Loop扮演Kubernetes控制器的角色，确保期望与实际的运行状态是一致的。
使用控制器模式，与Kubernetes里API对象的“增、删、改、查”进行协作，进而完成用户业务逻辑的编写过程，这就是声明式API对象的编程范式，即"Kubernetes编程范式"。


使用Operator SDK生成go项目框架

operator-sdk init --domain=example.com --repo=github.com/example-inc/vpc-operator

为刚才生成代码添加自定义API

operator-sdk create api --group network --version v1 --kind Vpc --resource=true --controller=true
添加自定义控制器。

operator-sdk add controller --api-version=network.example.com/v1 --kind=Vpc

修改api/v1/xxxx_types.go 文件里面的自定义的字段；

// VpcSpec defines the desired state of Vpc


type VpcSpec struct {

	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Foo is an example field of Vpc. Edit Vpc_types.go to remove/update
    
	VpcName   string `json:"vpcname,omitempty"`
	CidrBlock string `json:"cidrblock,omitempty"`
	SubnetNum int64  `json:"subnetnum,omitempty"`
}

// VpcStatus defines the observed state of Vpc
type VpcStatus struct {

	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	PodNames  []string `json:"podNames,omitempty"`
	SubnetNum int64    `json:"subnetnum,omitempty"`
}

监听资源
核心调度函数就是Reconcile

request就是发生变动的资源信息，主要就是namespace和name。

因为我们之前监听并加入队列的一定是Demo资源，所以我们直接利用k8s客户端Get获取发生变动的Demo对象。

func (r *ReconcileImoocPod) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	
    reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling ImoocPod")
	// Fetch the ImoocPod instance
	instance := &k8sv1alpha1.ImoocPod{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
    
	lbls := labels.Set{
		"app": instance.Name,
	}
	existingPods := &corev1.PodList{}
        
         //需要自己手动添加的地方
	 
	//1:获取name对应的所以的pod列表
	err = r.client.List(context.TODO(), existingPods, &client.ListOptions{
		Namespace:     request.Namespace,
		LabelSelector: labels.SelectorFromSet(lbls),
	})
	if err != nil {
		reqLogger.Error(err, "取已经存在的pod失败")
		return reconcile.Result{}, err
	}
 
	//2:取到pod列表中的pod name
	var existingPodNames []string
	for _, pod := range existingPods.Items {
		if pod.GetObjectMeta().GetDeletionTimestamp() != nil {
			continue
		}
 
		if pod.Status.Phase == corev1.PodPending || pod.Status.Phase == corev1.PodRunning {
			existingPodNames = append(existingPodNames, pod.GetObjectMeta().GetName())
		}
 
	}
 
	//3:update pod.status !=运行中的status
	//比较 DeepEqual
	status := k8sv1alpha1.ImoocPodStatus{ //期望的status
		PodNames: existingPodNames,
		Replicas: len(existingPodNames),
	}
	if !reflect.DeepEqual(instance.Status, status) {
		instance.Status = status //把期望状态给运行态
		err = r.client.Status().Update(context.TODO(), instance)
		if err != nil {
			reqLogger.Error(err, "更新pod的状态失败")
			return reconcile.Result{}, err
		}
	}
 
	//4:len(pod)>运行中的len(pod.replace)，期望值小，需要scale down
	if len(existingPodNames) > instance.Spec.Replicas {
		//delete
		reqLogger.Info("正在删除Pod,当前的podnames和期望的Replicas:", existingPodNames, instance.Spec.Replicas)
		pod := existingPods.Items[0]
		err := r.client.Delete(context.TODO(), &pod)
		if err != nil {
			reqLogger.Error(err, "删除pod失败")
			return reconcile.Result{}, err
		}
	}
 
	//5:len(pod)<运行中的len(pod.replace)，期望值大，需要scale up create
	if len(existingPodNames) < instance.Spec.Replicas {
		//create
		reqLogger.Info("正在创建Pod,当前的podnames和期望的Replicas:", existingPodNames, instance.Spec.Replicas)
		// Define a new Pod object
		pod := newPodForCR(instance)
 
		// Set ImoocPod instance as the owner and controller
		if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
			return reconcile.Result{}, err
		}
 
		err = r.client.Create(context.TODO(), pod)
		if err != nil {
			reqLogger.Error(err, "创建pod失败")
			return reconcile.Result{}, err
		}
	}
 
	 Define a new Pod object
	//pod := newPodForCR(instance)
	//
	 Set ImoocPod instance as the owner and controller
	//if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
	//	return reconcile.Result{}, err
	//}
	//
	 Check if this Pod already exists
	//found := &corev1.Pod{}
	//err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
	//if err != nil && errors.IsNotFound(err) {
	//	reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
	//	err = r.client.Create(context.TODO(), pod)
	//	if err != nil {
	//		return reconcile.Result{}, err
	//	}
	//
	//	// Pod created successfully - don't requeue
	//	return reconcile.Result{}, nil
	//} else if err != nil {
	//	return reconcile.Result{}, err
	//}
	//
	 Pod already exists - don't requeue
	//reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	//return reconcile.Result{}, nil
	return reconcile.Result{Requeue: true}, nil
}


运行controller有两种方法

可以在本地直接运行controller，也可以打包到k8s运行。

本地运行controller
在本地运行controller直接go run就可以了
export WATCH_NAMESPACE=default
go run cmd/manager/main.go

打包提交到k8s运行
如果我们controller完成，我们可以将其打包放到k8s上运行
