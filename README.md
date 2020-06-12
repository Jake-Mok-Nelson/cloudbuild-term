# cloudbuild-term
Review Google Cloudbuilds in the terminal

**Warning** This is an MVP as far as tools go. There's not much in the way of testing or mocking and it needs work. Use at your own risk.

## Why?

Who wants to use a GUI to track down cloudbuilds when there could be many that are running concurrently across multiple projects?

## How?

Ensure you are authorised to use Google APS in the terminal with `gcloud auth login` and `gcloud auth application-default login`

It is allowing application-default credentials to use your login that provides access to the application.
