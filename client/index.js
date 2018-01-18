const teller = require('@autopilot/teller');
const client = teller.impl({})({
  NetworkLayer: {
    // Because we haven't implemented acks in Banks yet. This won't be
    // needed in the near future.
    requestAcknowledgementTimeout: 4000,
  },
});
client.NetworkLayer.open((err, result) => {
  if (err) {
    outputError(err);
    return
  }
  client.Transport.connect((err, result) => {
    if (err) {
      outputError(err);;
      return
    }

    process.on('beforeExit', (code) => {
      client.NetworkLayer.close();
    });

    console.log("CONNECTION ID", result.body.connectionID);
    generateApiKey("my-role");
    enableApiKey("my-api-key");
    disableApiKey("my-api-key");
    getApiKey("my-role");
  });
});

function generateApiKey(role) {
  const params = {
    dataScope: "taranaki-user-1",
    resourceName: "ApiKeys",
    messageName: "Generate",
    body: {
      role,
    },
  };
  client.Transport.requestResponse(params, (err, result) => {
    if (err) {
      outputError(err);
      return
    }
    console.log("\n\nApiKeys.Generate response:\n", result.body);
  });
}

function enableApiKey(token) {
  const params = {
    dataScope: "taranaki-user-1",
    resourceName: "ApiKeys",
    messageName: "Enable",
    body: {
      token,
    },
  };
  client.Transport.requestResponse(params, (err, result) => {
    if (err) {
      outputError(err);
      return
    }
    console.log("\n\nApiKeys.Enable response:\n", result.body);
  });
}

function disableApiKey(token) {
  const params = {
    dataScope: "taranaki-user-1",
    resourceName: "ApiKeys",
    messageName: "Disable",
    body: {
      token,
    },
  };
  client.Transport.requestResponse(params, (err, result) => {
    if (err) {
      outputError(err);
      return
    }
    console.log("\n\nApiKeys.Disable response:\n", result.body);
  });
}

function getApiKey(role) {
  const params = {
    dataScope: "taranaki-user-1",
    resourceName: "ApiKeys",
    messageName: "Get",
    body: {
      role,
    },
  };
  client.Transport.requestResponse(params, (err, result) => {
    if (err) {
      outputError(err);
      return
    }
    console.log("\n\nApiKeys.Get response:\n", result.body);
  });
}

// output a Teller error in a human-readable format
function outputError(err) {
  if (err.details && err.details.responseHeaderErrors) {
    const errs = err.details.responseHeaderErrors;

    if (errs.length > 1) {
      for (const e of errs) {
        console.log(`\t${e.message} (Error code ${e.code})`);
      }
    } else {
      console.log(`\nError: ${errs[0].message} (Error code ${errs[0].code})`);
    }

  } else {
    console.log("\nError:", err.toString());
  }

  process.exitCode = 1;
}

// safely round floating point numbers to a specific number of decimal places.
function round(value, decimals) {
  return Number(Math.round(value+'e'+decimals)+'e-'+decimals);
}
