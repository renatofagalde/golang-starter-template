
bibliotecas
zap
gingonic
gorm
github.com/go-resty/resty/v2
get_env
copier


GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o bootstrap ./cmd/lambda_api

zip -j deployment.zip bootstrap

sam local invoke --event deployments/events/list_notes.json --template deployments/lambda/template.yaml

opcao menos -f para ficar em background
ssh -i .ssh/key.pem -f -N -L 5432:xxxxx-1-instance-1.cvwyig4oy864.us-east-1.rds.amazonaws.com:5432 ec2-user@ec2-34-xxx-xxx-xxx.compute-1.amazonaws.com -v


# Para ambiente local (usa o banco PostgreSQL em contêiner)
make local

# Para ambiente de desenvolvimento (conecta ao banco na AWS)
make dev

# Para ambiente de produção (conecta ao banco na AWS)
make prod

# Para subir toda a aplicação local (banco + migrations + app)
make up-local

# Para ver a ajuda
make help