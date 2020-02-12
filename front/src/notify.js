function notify (payload) {
  window.eventBus.$emit('vtNotify', {
    body: payload.body,
    /**
     * types:
     * ~ success
     * ~ error
     * ~ alert
     * ~ info
     */
    type: payload.type,
    canTimeout: true,
    duration: 3000
  })
}

export { notify }
