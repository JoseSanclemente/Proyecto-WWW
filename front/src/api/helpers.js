let helpers = {
  generateConfig: function (extraHeaders) {
    let basicConfig = {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      }
    }

    if (extraHeaders === null || extraHeaders === undefined) {
      return basicConfig
    }

    Object.keys(extraHeaders).forEach(function (key) {
      basicConfig.headers[key] = extraHeaders[key]
    })

    return basicConfig
  }
}

export { helpers }
