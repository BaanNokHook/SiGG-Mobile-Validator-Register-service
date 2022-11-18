template from: https://git.matador.ais.co.th/cronus/irm/devops/gitlab-init-template.git

device-register
741W36Cz&jtu

helm template ./helm -f helm/values.yaml --name-template mobile-validator-register-service --set deployment.image="registry.gitlab.com/nextdb-project/digital-reality-foundation/mobile-validator/validator-register/mobile-validator-register-service:latest" -n transaction-gateway-be > helm/k8s/deployment.yml

helm install ./helm -f helm/values.yaml --name-template mobile-validator-register-service --set deployment.image="registry.gitlab.com/nextdb-project/digital-reality-foundation/mobile-validator/validator-register/mobile-validator-register-service:latest" -n transaction-gateway-be

kubectl create secret docker-registry gitlab-cr --docker-server=registry.gitlab.com --docker-username=gitlab+deploy-token-1458146 --docker-password=BRHQZxvUyrVyLWnBNP8\_ --docker-email=p.naksinn@gmail.com -n mobile-register-loadtest

helm upgrade --set deployment.image="registry.gitlab.com/nextdb-project/digital-reality-foundation/mobile-validator/validator-register/mobile-validator-register-service:develop" -f helm/mobile-validator-register-dev-values.yaml loadtest-1 ./helm -n mobile-register-loadtest --install
