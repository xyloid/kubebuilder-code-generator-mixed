#!/usr/bin/env bash

kubectl delete -f config/samples/webapp_v1_guestbook.yaml
kubectl delete -f config/crd/bases/webapp.example.com_guestbooks.yaml
