build-k8s:
	@docker build -t plantdex/plants:prod --target prod plants && \
	docker build -t plantdex/users:prod --target prod users && \
	docker build -t plantdex/web:prod --target prod web

start:
	@docker compose -p plantdex up --build -d

start-%:
	@docker compose -p plantdex up --build $*

start-k8s: build-k8s
	@cd .kubernetes && \
	kubectl create namespace plantdex; \
	kubectl config set-context --current --namespace=plantdex && \
	kubectl apply -f users-secret.yaml && \
	kubectl apply -f plants-secret.yaml && \
	kubectl apply -f users-configmap.yaml && \
	kubectl apply -f plants-configmap.yaml && \
	kubectl apply -f web-configmap.yaml && \
	kubectl apply -f users.yaml && \
	kubectl apply -f plants.yaml && \
	kubectl apply -f web.yaml && \
	kubectl apply -f ingress.yaml && \
	cd ..

stop:
	@docker compose -p plantdex down

stop-k8s:
	@kubectl delete ns plantdex
