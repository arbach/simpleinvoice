FROM ubuntu

ENV DEBIAN_FRONTEND=noninteractive
RUN apt update && apt upgrade  -y &&  \
    apt install -y apache2 apache2-utils  &&  \
    a2enmod headers &&  \
    a2enmod proxy &&  \
    a2enmod proxy_http && \
    a2enmod proxy_balancer && \
    a2enmod lbmethod_byrequests && \
    service apache2 start
 
COPY . /var/www/html/
COPY 000-default.conf /etc/apache2/sites-enabled/000-default.conf


EXPOSE 80 
CMD ["apache2ctl", "-D", "FOREGROUND"]
