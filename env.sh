#! /bin/sh
# command to run go cmd file
# Should not have this file pushed to git but
# for demonstration purpose. This might be not 
# the best way to store database credentials...
# It might be better to set permission using chmod 700
# on this file, i guess?
alias gorun='GOINTVPR_DB_USER="admin" GOINTVPR_DB_PASSWORD="admin" go run cmd/main.go'
