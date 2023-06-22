pipeline {
    agent {
            label 'jenkins-jnlp-slave'
    }
    stages {
        stage('Clone Repo') {
            steps { 
                checkout scm
            }
        }
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
            container('test') {
                script {                 
                    sh '''
                    /kaniko/executor --dockerfile `pwd`/Dockerfile --context `pwd` --destination=switch2mdock/elastic:99               
                    '''
                }          
            }           
        }         
    }

    }    
}
    
        
    

