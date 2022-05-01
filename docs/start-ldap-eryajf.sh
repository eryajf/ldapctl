#!/bin/bash -e

SERVICE=ldap-service-eryajf
HOST_NAME=ldap-server-eryajf
LDAP_DOMAIN=eryajf.net
LDAP_DC=eryajf
LDAP_DC_ORG=net
NETWORK_ADAPTER=eth0
PASSWORD=123465
OPENLDAP="1.5.0"
PHPLDAPADMIN="0.9.0"
HTTPS_PORT=8089
OPENLDAP_PORT=390

docker run \
    -p ${OPENLDAP_PORT}:389 \
    --name ${SERVICE} \
    --hostname ${HOST_NAME} \
    --env LDAP_ORGANISATION="Eyajf-Group" \
    --env LDAP_DOMAIN=${LDAP_DOMAIN} \
    --env LDAP_ADMIN_PASSWORD=${PASSWORD} \
    --detach osixia/openldap:${OPENLDAP}

docker run \
    -p ${HTTPS_PORT}:80 \
    --name ${SERVICE}-admin \
    --hostname ${HOST_NAME}-admin \
    --link ${SERVICE}:${HOST_NAME} \
    --env PHPLDAPADMIN_LDAP_HOSTS=${HOST_NAME} \
    --env PHPLDAPADMIN_HTTPS=false \
    --detach \
    osixia/phpldapadmin:${PHPLDAPADMIN}

sleep 1
echo "-----------------------------------"
PHPLDAP_IP=$(docker inspect -f "{{ .NetworkSettings.IPAddress }}" ${SERVICE})
docker exec ${SERVICE} ldapsearch -x -H ldap://${PHPLDAP_IP}:389 -b "dc=${LDAP_DC},dc=${LDAP_DC_ORG}" -D "cn=admin,dc=${LDAP_DC},dc=${LDAP_DC_ORG}" -w ${PASSWORD}
echo "-----------------------------------"
# If it is not debugged locally, replace it with host IP here.
PUB_IP="localhost"
echo "Go to: https://${PUB_IP}:${HTTPS_PORT}"
echo "Login DN: cn=admin,dc=${LDAP_DC},dc=${LDAP_DC_ORG}"
echo "Password: ${PASSWORD}"