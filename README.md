# HTTP-GPIO Server

Server implementation of the HTTP-GPIO API specification, an open standard
that describes a way of interfacing with an output module via a RESTful
HTTP API server.

## Methods

**Please note:** outputs are 1-based

Turn on an output:

    PUT /outputs/2
    <- {"index": 2, "position": 1 }

Turn off an output:

    DELETE /outputs/2
    <- {"index": 2, "position": 0 }

Get the position of an output:

    GET /outputs/2
    <- {"index": 2, "position": 1 }

Turn off all outputs:

    DELETE /outputs
    <- [{"index": 1, "position": 0 }, {"index": 2, "position": 0 }]

Turn on all outputs:

    PUT /outputs
    <- [{"index": 1, "position": 1 }, {"index": 2, "position": 1 }]
