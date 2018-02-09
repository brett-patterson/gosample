# gosample

A sample Go web application deployable to Google's Flexible App Engine. Uses the [Echo](https://echo.labstack.com/) library
for the web server for ease of routing requests, receiving and sending JSON, logging requests, and handling a
production-level volume of requests.

Also provides a sample [request middleware](https://echo.labstack.com/middleware) that would be used
for authentication and authorization.

To deploy, install the [Google Cloud SDK](https://cloud.google.com/sdk/docs/) and then run:

```bash
gcloud app deploy
```

This app is already deployed to https://gleaming-plate-194700.appspot.com/.

Try the following endpoints:
* https://gleaming-plate-194700.appspot.com/
    * The home endpoint that does not go through the security middleware
* https://gleaming-plate-194700.appspot.com/equipment
    * The equipment endpoint that passes through the security middleware and returns sample equipment JSON
* https://gleaming-plate-194700.appspot.com/equipment?pretty
    * A convenience provided by the Echo library on any JSON endpoint to format the JSON nicely for humans
