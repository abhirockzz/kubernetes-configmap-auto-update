## Kubernetes ConfigMap auto update

Using environment variables in your application (`Deployment`) via `ConfigMap` poses a challenge - how will your app uptake the new values in case the `ConfigMap` gets updated? You can obviously recreate the deployment - but this is undesirable in most cases.

However, if you load `ConfigMap` contents as a volume, you get the auto update capability for *free* i.e. Kubernetes ensures that the volume contents (files containing the config value) are refresh if the `ConfigMap` gets updated.