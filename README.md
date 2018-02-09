# gosample

A sample Go web application deployable to Google's Flexible App Engine. Uses the [echo](https://echo.labstack.com/) library
for the web server for ease of routing requests, receiving and sending JSON, logging requests, and handling a
production-level volume of requests.

Also provides a sample [request middleware](https://echo.labstack.com/middleware) that would be used
for authentication and authorization.

To deploy, install the [Google Cloud SDK](https://cloud.google.com/sdk/docs/) and then run:

```bash
gcloud app deploy
```
