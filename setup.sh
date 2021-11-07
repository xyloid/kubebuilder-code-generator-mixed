#!/usr/bin/env bash

kubectl apply -f config/crd/bases/webapp.example.com_guestbooks.yaml
kubectl apply -f config/samples/webapp_v1_guestbook.yaml