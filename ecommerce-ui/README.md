# Ecommerce UI React App

This is the frontend component of the <em>Spring boot </em> based <strong>ecommerce </strong> web service. Ensure that this service is up either in docker or kubernetes environment before starting the frontend service.

## Building docker image
This project already has a Dockerfile in it. Just build it either through roost GUI or by the running the following command in your terminal at the root of this project <br> <ul>
<li>
    <code>docker build -t ecommerce-ui:v1 .</code>
    <br> This will create a image named <strong>ecommerce-ui:v1</strong>
</li>
</ul>

## Using this image in pod or deployment yaml
This image expects an environment variable called <strong>REACT_APP_API_URL</strong> in its npm run command i.e. <br>
<code>REACT_APP_API_URL=roost-utility:8080 npm start</code> should be its startup command if the backend service is running just as a docker container or <br>
<code>REACT_APP_API_URL=roost-controlplane: NODE_PORT npm start</code> where <strong>NODE_PORT</strong> will be the nodePort of the backend service running in the kubernetes environment. If this environment variable is not mentioned the react app will try to fetch the products from <em>localhost:8080</em> which will not work in these environments.
