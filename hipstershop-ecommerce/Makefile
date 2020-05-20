EMAIL_SERVICE_IMG=emailservice
PRODUCT_CATALOG_SERVICE_IMG=productcatalogservice
RECOMMENDATION_SERVICE_IMG=recommendationservice
SHIPPING_SERVICE_IMG=shippingservice
CHECKOUTSERVICE_IMG=checkoutservice
PAYMENT_SERVICE_IMG=paymentservice
CURRENCY_SERVICE_IMG=currencyservice
CARTSERVICE_IMG=cartservice
FRONTEND_IMG=frontend
LOAD_GENERATOR_IMG=loadgenerator
ADDSERVICE_IMG=adservice

IMAGE_TAG=v1

GRPC_HEALTH_PROBE=grpc_health_probe-linux-amd64
GRPC_HEALTH_PROBE_PATH=../tool/${GRPC_HEALTH_PROBE}

HOSTNAME := $(shell hostname -f)

GO_SERVICES = src/${PRODUCT_CATALOG_SERVICE_IMG} src/${SHIPPING_SERVICE_IMG} src/${CHECKOUTSERVICE_IMG} src/${FRONTEND_IMG}
MICROSERVICES = ${EMAIL_SERVICE_IMG} \
				${PRODUCT_CATALOG_SERVICE_IMG} \
				${RECOMMENDATION_SERVICE_IMG}	\
				${SHIPPING_SERVICE_IMG}	\
				${CHECKOUTSERVICE_IMG}	\
				${PAYMENT_SERVICE_IMG}	\
				${CURRENCY_SERVICE_IMG}	\
				${CARTSERVICE_IMG}	\
				${FRONTEND_IMG}	\
				${LOAD_GENERATOR_IMG}	\
				${ADDSERVICE_IMG}

SERVICE_GRPC_BIN = $(foreach microservice, ${MICROSERVICES}, "src/${microservice}"/${GRPC_HEALTH_PROBE})

.PHONY: all
all: clean prerequisites dockerise deploy
	$(MAKE) clean_bin

.PHONY: prerequisites
prerequisites:
	for service in ${MICROSERVICES}; do \
		if ( [ $${service} != ${FRONTEND_IMG} ] && [ $${service} != ${LOAD_GENERATOR_IMG} ] ); then \
			cp ${GRPC_HEALTH_PROBE_PATH} src/$${service}; \
		fi; \
	done

.PHONY: build
build:
	# go build -o src/checkoutservice/bin/checkoutservice -i -mod vendor -modfile go.mod -gcflags='-N -l' github.com/roost-io/roost-example/hipstershop-ecommerce/src/checkoutservice
	for d in $(GO_SERVICES); do echo "<==> Target $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done

.PHONY: dockerise
dockerise:
	for service in ${MICROSERVICES}; do \
		docker build -f src/$${service}/Dockerfile -t $${service}:${IMAGE_TAG} ./src/$${service} ; \
	done

.PHONY: deploy
deploy:
	kubectl apply -f kubernetes-manifests/zke-deployment.yaml

.PHONY: clean
clean: clean_bin 
	kubectl delete -f kubernetes-manifests/zke-deployment.yaml --now >/dev/null 2>&1 || true

.SILENT:
clean_bin:
	# -rm -f ${SERVICE_GRPC_BIN}
	#for service in ${MICROSERVICES}; do \
		echo $${service}
		# if [ -f "src/$${service}/$${service}" ]; then \
		# 	rm -f src/$${service}/$${service}; \
		# fi; \

		# if [ -f "src/$${service}/${GRPC_HEALTH_PROBE}" ]; then \
		# 	rm -f src/$${service}/${GRPC_HEALTH_PROBE}; \
		# fi; \
	
#	done

test:
	@(for grpc_bin in ${SERVICE_GRPC_BIN}; do \
		echo $${grpc_bin}; done;	\
	)