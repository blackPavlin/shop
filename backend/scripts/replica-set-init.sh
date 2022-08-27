#!/bin/bash

mongo <<EOF
var config = {
    "_id" : "rsmongo",
    "members" : [
        { "_id" : 0, "host" : "mongo:27017" }
    ]
}
rs.initiate(config, { force: true });
EOF
