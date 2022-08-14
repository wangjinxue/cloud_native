openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj \
'/O=cloudnative Inc./CN=*.51.cafe' -keyout 51.cafe.key -out 51.cafe.crt
kubectl create -n istio-system secret tls wildcard-credential --key=51.cafe.key --cert=51.cafe.crt