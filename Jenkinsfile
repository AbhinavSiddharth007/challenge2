// pipeline {
//   agent any

//   stages {
//     stage('Build') {
//       steps {
//         sh 'go build -o main .'
//       }
//     }

//     stage('Deploy') {
//       steps {
//         sh '''
//           set -eu

//           ssh target 'sudo systemctl stop myapp || true'
//           ssh target 'sudo mkdir -p /opt/myapp && sudo chown "$(whoami)" /opt/myapp && sudo chmod 755 /opt/myapp'

//           scp main target:/opt/myapp/main
//           ssh target 'sudo chmod +x /opt/myapp/main'

//           scp myapp.service target:/tmp/myapp.service
//           ssh target 'sudo install -m 644 /tmp/myapp.service /etc/systemd/system/myapp.service && sudo rm -f /tmp/myapp.service'

//           ssh target 'id -u myapp >/dev/null 2>&1 || sudo useradd --system --no-create-home --shell /usr/sbin/nologin myapp'
//           ssh target 'sudo systemctl daemon-reload'
//           ssh target 'sudo systemctl enable myapp'
//           ssh target 'sudo systemctl restart myapp'
//         '''
//       }
//     }

//     stage('Health Check') {
//       steps {
//         sh '''
//           set -eu

//           i=1
//           while [ "$i" -le 10 ]; do
//             if curl -fsS http://target:4444/; then
//               exit 0
//             fi
//             i=$((i + 1))
//             sleep 2
//           done

//           exit 1
//         '''
//       }
//     }
//   }
// }

pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh 'go build -o main .'
      }
    }
    stage('Deploy') {
      steps {
        sh '''
          set -eu

          ssh -o StrictHostKeyChecking=accept-new root@target 'sudo systemctl stop myapp || true'
          ssh -o StrictHostKeyChecking=accept-new root@target 'sudo mkdir -p /opt/myapp && sudo chown "$(whoami)" /opt/myapp && sudo chmod 755 /opt/myapp'

          scp -o StrictHostKeyChecking=accept-new main root@target:/opt/myapp/main
          ssh -o StrictHostKeyChecking=accept-new root@target 'sudo chmod +x /opt/myapp/main'

          scp -o StrictHostKeyChecking=accept-new myapp.service root@target:/tmp/myapp.service
          ssh -o StrictHostKeyChecking=accept-new root@target 'sudo install -m 644 /tmp/myapp.service /etc/systemd/system/myapp.service && sudo rm -f /tmp/myapp.service'

          ssh -o StrictHostKeyChecking=accept-new root@target 'id -u myapp >/dev/null 2>&1 || sudo useradd --system --no-create-home --shell /usr/sbin/nologin myapp'
          ssh -o StrictHostKeyChecking=accept-new root@target 'sudo systemctl daemon-reload'
          ssh -o StrictHostKeyChecking=accept-new root@target 'sudo systemctl enable myapp'
          ssh -o StrictHostKeyChecking=accept-new root@target 'sudo systemctl restart myapp'
        '''
      }
    }
    stage('Health Check') {
      steps {
        sh '''
          set -eu

          i=1
          while [ "$i" -le 10 ]; do
            if curl -fsS http://target:4444/; then
              echo
              echo "App is serving."
              exit 0
            fi
            i=$((i + 1))
            sleep 2
          done

          echo "App never came up after 10 tries" >&2
          exit 1
        '''
      }
    }
  }
}
