pipeline{
    agent{
        kubernetes{
            defaultContainer 'golang'
            yaml '''
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: golang
    image: golang:1.20.3
    command:
    - sleep
    args:
    - 99d
  - name: kaniko
    image: gcr.io/kaniko-project/executor:debug
    imagePullPolicy: Always
    command:
    - sleep
    args:
    - 9999999
    volumeMounts:
    - mountPath: /kaniko/.docker
      name: kaniko-secret
  volumes:
    - name: kaniko-secret
      secret:
        secretName: regcred
        items:
          - key: .dockerconfigjson
            path: config.json
            '''
            }
    }
    stages{
        stage('Test'){
            steps{
                container('golang'){
                //sh 'git clone https://github.com/web2543/go-ci.git .'
                sh 'go test -v'
                }
            }
        }
        stage('Build'){
            steps{
                container('golang'){
                echo "Running $BUILD_ID,$BUILD_NUMBER"
                sh 'go build -buildvcs=false'
                sh 'ls'
                }  
            }
        }
        stage('Container'){
            steps {
                    container(name: 'kaniko', shell: '/busybox/sh') {
                        sh '''#!/busybox/sh
                            /kaniko/executor --context `pwd` --destination tanaxw/tanax:0.0.$BUILD_NUMBER
                            '''
                    }
                }
            }
        }
}
