MICROSERVICE_NAME=productcatalogservice

build:
	go build -o ${MICROSERVICE_NAME} -i -mod vendor -modfile ../../go.mod -gcflags='-N -l' github.com/roost-io/roost-example/hipstershop-ecommerce/src/${MICROSERVICE_NAME}

clean_bin:
	-rm -f