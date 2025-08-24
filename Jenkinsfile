pipeline 
  agent any

  environment {
    IMAGE_NAME = 'user-service'
    DOCKER_CREDENTIALS = credentials('docker-credential')
    GITH    stage('Checkout Code') {
      steps {
        script {
          echo "====== CHECKOUT DEBUG ======"
          def repoUrl = 'https://github.com/Mini-Soccer-Project/user-service.git'
          echo "Repository URL: ${repoUrl}"
          echo "Target Branch: ${env.TARGET_BRANCH}"
          echo "Credentials ID: github-credential"
          
          try {
            echo "Starting checkout process..."
            checkout([$class: 'GitSCM',
              branches: [
                [name: "*/${env.TARGET_BRANCH}"]
              ],
              userRemoteConfigs: [
                [url: repoUrl, credentialsId: 'github-credential']
              ]
            ])
            
            echo "Checkout successful! Listing workspace contents:"
            sh '''
              echo "Current directory: $(pwd)"
              echo "Directory contents:"
              ls -lah
              echo "Git status after checkout:"
              git status
              echo "Current branch:"
              git branch
              echo "Git log (last 3 commits):"
              git log --oneline -3
            '''
          } catch (Exception e) {
            echo "Checkout failed with error: ${e.getMessage()}"
            echo "Stack trace: ${e.getStackTrace()}"
            throw e
          }
        }
      }
    }redentials('github-credential')
    SSH_KEY = credentials('ssh-key')
    HOST = credentials('host')
    USERNAME = credentials('username')
    CONSUL_HTTP_URL = credentials('consul-http-url')
    CONSUL_HTTP_TOKEN = credentials('consul-http-token')
    CONSUL_WATCH_INTERVAL_SECONDS = 60
  }

  stages {
    stage('Environment Debug') {
      steps {
        script {
          echo "====== JENKINS ENVIRONMENT DEBUG ======"
          echo "Build Number: ${currentBuild.number}"
          echo "Build ID: ${env.BUILD_ID}"
          echo "Build URL: ${env.BUILD_URL}"
          echo "Jenkins URL: ${env.JENKINS_URL}"
          echo "Job Name: ${env.JOB_NAME}"
          echo "Build Tag: ${env.BUILD_TAG}"
          echo "Node Name: ${env.NODE_NAME}"
          echo "Workspace: ${env.WORKSPACE}"
          echo "GIT_BRANCH: ${env.GIT_BRANCH}"
          echo "GIT_COMMIT: ${env.GIT_COMMIT}"
          echo "GIT_URL: ${env.GIT_URL}"
          
          echo "====== SYSTEM ENVIRONMENT DEBUG ======"
          sh '''
            echo "Current user: $(whoami)"
            echo "Current directory: $(pwd)"
            echo "HOME directory: $HOME"
            echo "PATH: $PATH"
            echo "Shell: $SHELL"
            echo "OS Info: $(uname -a)"
            echo "Available disk space:"
            df -h
            echo "Memory info:"
            free -h || echo "free command not available"
            echo "Environment variables:"
            env | sort
          '''
          
          echo "====== GIT DEBUG ======"
          sh '''
            echo "Checking Git installation:"
            which git || echo "Git not found in PATH"
            git --version || echo "Git command failed"
            ls -la /usr/bin/git || echo "Git not in /usr/bin"
            ls -la /usr/local/bin/git || echo "Git not in /usr/local/bin"
            
            echo "Git configuration:"
            git config --list || echo "Git config failed"
            
            echo "Current Git status:"
            git status || echo "Git status failed"
            git remote -v || echo "Git remote failed"
            git branch -a || echo "Git branch failed"
          '''
          
          echo "====== DOCKER DEBUG ======"
          sh '''
            echo "Docker version:"
            docker --version || echo "Docker not available"
            docker info || echo "Docker info failed"
            docker ps || echo "Docker ps failed"
          '''
        }
      }
    }

    stage('Check Commit Message') {
      steps {
        script {
          echo "====== COMMIT MESSAGE CHECK DEBUG ======"
          try {
            def commitMessage = sh(
              script: "git log -1 --pretty=%B",
              returnStdout: true
            ).trim()

            echo "Commit Message: ${commitMessage}"
            echo "Commit Hash: ${sh(script: 'git rev-parse HEAD', returnStdout: true).trim()}"
            echo "Author: ${sh(script: 'git log -1 --pretty=%an', returnStdout: true).trim()}"
            echo "Date: ${sh(script: 'git log -1 --pretty=%ad', returnStdout: true).trim()}"
            
            if (commitMessage.contains("[skip ci]")) {
              echo "Skipping pipeline due to [skip ci] tag in commit message."
              currentBuild.result = 'ABORTED'
              currentBuild.delete()
              return
            }

            echo "Pipeline will continue. No [skip ci] tag found in commit message."
          } catch (Exception e) {
            echo "Error checking commit message: ${e.getMessage()}"
            echo "Stack trace: ${e.getStackTrace()}"
            // Continue pipeline despite error
          }
        }
      }
    }

    stage('Set Target Branch') {
      steps {
        script {
          echo "====== BRANCH SETUP DEBUG ======"
          echo "GIT_BRANCH: ${env.GIT_BRANCH}"
          echo "BRANCH_NAME: ${env.BRANCH_NAME}"
          echo "CHANGE_BRANCH: ${env.CHANGE_BRANCH}"
          
          // Debug all environment variables related to Git
          sh '''
            echo "All Git-related environment variables:"
            env | grep -i git || echo "No Git environment variables found"
          '''
          
          if (env.GIT_BRANCH == 'origin/main') {
            env.TARGET_BRANCH = 'main'
            echo "Setting TARGET_BRANCH to 'main'"
          } else if (env.GIT_BRANCH == 'origin/development') {
            env.TARGET_BRANCH = 'development'
            echo "Setting TARGET_BRANCH to 'development'"
          } else {
            env.TARGET_BRANCH = 'main' // Default fallback
            echo "Using default TARGET_BRANCH 'main' for branch: ${env.GIT_BRANCH}"
          }

          echo "Final TARGET_BRANCH: ${env.TARGET_BRANCH}"
        }
      }
    }

    stage('Credentials Debug') {
      steps {
        script {
          echo "====== CREDENTIALS DEBUG ======"
          try {
            echo "Checking Docker credentials availability..."
            withCredentials([usernamePassword(credentialsId: 'docker-credential', passwordVariable: 'DOCKER_PASSWORD', usernameVariable: 'DOCKER_USERNAME')]) {
              echo "Docker Username: ${DOCKER_USERNAME}"
              echo "Docker Password: [HIDDEN]"
              echo "Docker credential length: ${DOCKER_PASSWORD.length()}"
            }
          } catch (Exception e) {
            echo "Docker credentials error: ${e.getMessage()}"
          }
          
          try {
            echo "Checking GitHub credentials availability..."
            withCredentials([usernamePassword(credentialsId: 'github-credential', passwordVariable: 'GITHUB_PASSWORD', usernameVariable: 'GITHUB_USERNAME')]) {
              echo "GitHub Username: ${GITHUB_USERNAME}"
              echo "GitHub Password: [HIDDEN]"
              echo "GitHub credential length: ${GITHUB_PASSWORD.length()}"
            }
          } catch (Exception e) {
            echo "GitHub credentials error: ${e.getMessage()}"
          }
          
          try {
            echo "Checking SSH key availability..."
            withCredentials([sshUserPrivateKey(credentialsId: 'ssh-key', keyFileVariable: 'SSH_KEY')]) {
              echo "SSH Key file: ${SSH_KEY}"
              sh "ls -la ${SSH_KEY} || echo 'SSH key file not found'"
            }
          } catch (Exception e) {
            echo "SSH key error: ${e.getMessage()}"
          }
          
          try {
            echo "Checking Consul credentials..."
            echo "Consul URL: ${env.CONSUL_HTTP_URL}"
            echo "Consul Token: [HIDDEN]"
          } catch (Exception e) {
            echo "Consul credentials error: ${e.getMessage()}"
          }
        }
      }
    }

    stage('Checkout Code') {
      steps {
        script {
          def repoUrl = 'https://github.com/AndreaSuryatanaya/Mcs-user-service.git'

          checkout([$class: 'GitSCM',
            branches: [
              [name: "*/${env.TARGET_BRANCH}"]
            ],
            userRemoteConfigs: [
              [url: repoUrl, credentialsId: 'github-credential']
            ]
          ])

          sh 'ls -lah'
        }
      }
    }

    // stage('Login to Docker Hub') {
    //   steps {
    //     script {
    //       withCredentials([usernamePassword(credentialsId: 'docker-credential', passwordVariable: 'DOCKER_PASSWORD', usernameVariable: 'DOCKER_USERNAME')]) {
    //         sh """
    //         echo $DOCKER_PASSWORD | docker login -u $DOCKER_USERNAME --password-stdin
    //         """
    //       }
    //     }
    //   }
    // }

    // stage('Build and Push Docker Image') {
    //   steps {
    //     script {
    //       def runNumber = currentBuild.number
    //       sh "docker build -t ${DOCKER_CREDENTIALS_USR}/${IMAGE_NAME}:${runNumber} ."
    //       sh "docker push ${DOCKER_CREDENTIALS_USR}/${IMAGE_NAME}:${runNumber}"
    //     }
    //   }
    // }

    // stage('Update docker-compose.yaml') {
    //   steps {
    //     script {
    //       def runNumber = currentBuild.number
    //       sh "sed -i 's|image: ${DOCKER_CREDENTIALS_USR}/${IMAGE_NAME}:[0-9]\\+|image: ${DOCKER_CREDENTIALS_USR}/${IMAGE_NAME}:${runNumber}|' docker-compose.yaml"
    //     }
    //   }
    // }

    // stage('Commit and Push Changes') {
    //   steps {
    //     script {
    //       sh """
    //       git config --global user.name 'Jenkins CI'
    //       git config --global user.email 'jenkins@company.com'
    //       git remote set-url origin https://${GITHUB_CREDENTIALS_USR}:${GITHUB_CREDENTIALS_PSW}@github.com/Mini-Soccer-Project/user-service.git
    //       git add docker-compose.yaml
    //       git commit -m 'Update image version to ${TARGET_BRANCH}-${currentBuild.number} [skip ci]' || echo 'No changes to commit'
    //       git pull origin ${TARGET_BRANCH} --rebase
    //       git push origin HEAD:${TARGET_BRANCH}
    //       """
    //     }
    //   }
    // }

    // stage('Deploy to Remote Server') {
    //   steps {
    //     script {
    //       def targetDir = "/home/faisalilhami/mini-soccer-project/user-service"
    //       def sshCommandToServer = """
    //       ssh -o StrictHostKeyChecking=no -i ${SSH_KEY} ${USERNAME}@${HOST} '
    //         if [ -d "${targetDir}/.git" ]; then
    //             echo "Directory exists. Pulling latest changes."
    //             cd "${targetDir}"
    //             git pull origin "${TARGET_BRANCH}"
    //         else
    //             echo "Directory does not exist. Cloning repository."
    //             git clone -b "${TARGET_BRANCH}" git@github.com:Mini-Soccer-Project/user-service.git "${targetDir}"
    //             cd "${targetDir}"
    //         fi

    //         cp .env.example .env
    //         sed -i "s/^TIMEZONE=.*/TIMEZONE=Asia\\/Jakarta/" "${targetDir}/.env"
    //         sed -i "s/^CONSUL_HTTP_URL=.*/CONSUL_HTTP_URL=${CONSUL_HTTP_URL}/" "${targetDir}/.env"
    //         sed -i "s/^CONSUL_HTTP_PATH=.*/CONSUL_HTTP_PATH=backend\\/user-service/" "${targetDir}/.env"
    //         sed -i "s/^CONSUL_HTTP_TOKEN=.*/CONSUL_HTTP_TOKEN=${CONSUL_HTTP_TOKEN}/" "${targetDir}/.env"
    //         sed -i "s/^CONSUL_WATCH_INTERVAL_SECONDS=.*/CONSUL_WATCH_INTERVAL_SECONDS=${CONSUL_WATCH_INTERVAL_SECONDS}/" "${targetDir}/.env"
    //         sudo docker compose up -d --build --force-recreate
    //       '
    //       """
    //       sh sshCommandToServer
    //     }
    //   }
    // }
  }
  
  post {
    always {
      script {
        echo "====== POST-BUILD DEBUG ======"
        echo "Build Result: ${currentBuild.result}"
        echo "Build Duration: ${currentBuild.durationString}"
        echo "Build Description: ${currentBuild.description}"
        
        sh '''
          echo "Final workspace state:"
          ls -lah || echo "Cannot list workspace"
          echo "Disk space after build:"
          df -h || echo "Cannot check disk space"
          echo "Process list:"
          ps aux | head -10 || echo "Cannot list processes"
        '''
      }
    }
    
    success {
      script {
        echo "====== BUILD SUCCESS DEBUG ======"
        echo "Pipeline completed successfully!"
        echo "Build Number: ${currentBuild.number}"
        echo "Commit: ${env.GIT_COMMIT}"
      }
    }
    
    failure {
      script {
        echo "====== BUILD FAILURE DEBUG ======"
        echo "Pipeline failed!"
        echo "Build Number: ${currentBuild.number}"
        echo "Failed Stage: Check Jenkins console for details"
        
        sh '''
          echo "System state at failure:"
          echo "Last 50 lines of system log (if available):"
          tail -50 /var/log/syslog 2>/dev/null || echo "System log not accessible"
          echo "Jenkins log (if available):"
          tail -50 /var/log/jenkins/jenkins.log 2>/dev/null || echo "Jenkins log not accessible"
        '''
      }
    }
    
    unstable {
      script {
        echo "====== BUILD UNSTABLE DEBUG ======"
        echo "Pipeline completed with warnings!"
      }
    }
    
    aborted {
      script {
        echo "====== BUILD ABORTED DEBUG ======"
        echo "Pipeline was aborted!"
      }
    }
    
    cleanup {
      script {
        echo "====== CLEANUP DEBUG ======"
        sh '''
          echo "Cleaning up workspace..."
          echo "Current processes:"
          ps aux | grep -E "(docker|git|jenkins)" | grep -v grep || echo "No relevant processes found"
        '''
      }
    }
  }
