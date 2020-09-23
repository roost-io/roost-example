# Ecommerce web-app backend

## REST API exposed with Spring Boot

The backend consists of a H2 database containing catalogue of products

## Building the Docker image
Staying at the root directory (the directory where <em>pom.xml</em> file is located) of the project in a terminal execute the following command
<code>./mvnw spring-boot:build-image</code>
<br>
This will build a docker image for you by the name of <strong>ecommerce:v1</strong>


## Starting the container
<code>docker run -it -p8080:8080 ecommerce:v1</code>
<br>
The following command will start a container exposed at port 8080 with logs visible at your terminal

### The following APIs are exposed
<ul>
<li>Product list: <a>http://roost-utility:8080/api/products</a></li>
<li>Database console: <a>http://roost-utility:8080/h2-console</a> <br>
In the login page of the database enter the following values
<ol>
<li>Driver Class: <strong>org.h2.Driver</strong></li>
<li>
    JDBC URL: <a>jdbc:h2:mem:ecommercedb</a>
</li>
<li>Let the User Name be: <strong>sa</strong></li>
</ol>
</li>
The database has no password let the password field be empty, now you can hit connect. You should be able to see a table name <em>PRODUCT</em>
<br>
The table has a column called <strong>IMAGE_URL </strong>, it indicates the relative path of a product's image in the backend server
<li>Individual product images are also exposed: <a>http://roost-utility:8080/api/products/image/{IMAGE_URL}</a>
<br> You can get the IMAGE_URL of a product from the database as described above or from the product list endpoint. For eg one of the values of IMAGE_URL is <em>Watch</em>. Thus it's url will be <a>http://roost-utility:8080/api/products/image/Watch</a>
</li>
</ul>