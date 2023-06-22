pipeline {
    agent {
    kubernetes {
      yamlFile 'builder.yaml'
    }
    }
    stages {
    stage('Kaniko Build & Push Image') {
      steps {
        container('test') {
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
