Home-Reverse proxy
=================

# Goals

- Multiple hosts
- Configuration outside of binary

# Road map

- SSL Support
- Let's Encrypt auto-renewal

# Running

## Environment variables
- Create an environment variable for `ENV_KEY_LISTEN_ADDRESS`. The default value is: `127.0.0.1:8000`.
- Create an environment variable for `ENV_KEY_ROUTES`. The default value is: `localhost=http://team142.co.za` which is probably not helpful to you. You can have multiple routes setup in the format HOST=ROUTE,HOST2=ROUTE2


