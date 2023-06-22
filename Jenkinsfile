pipeline {
    agent {
    kubernetes {
      yamlFile 'builder.yaml'
    }
    stages {
        stage('Test environnement stage') {
            steps {
                echo 'test the first mvn package'
                sh '''
                    pwd
                    ls
                '''
            }
        }
        
    stage('Kaniko Build & Push Image') {
      steps {
        container('kaniko') {
          script {
            sh '''
            /kaniko/executor --dockerfile `pwd`/Dockerfile \
                             --context `pwd` \
                             --destination=switch2mdock/elastic:${BUILD_NUMBER}
            '''
          }
        }
      }
    }
  }
}
//cat << EOF > Dockerfile
//FROM alpine:latest
//RUN pwd
//EOF
    
        
    

