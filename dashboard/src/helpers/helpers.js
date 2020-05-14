export function getTransformers(substations) {
    let transformerList = []

    for (let i = 0; i < substations.length; i++) {
        let sub = substations[i]
        if (sub.transformers == null) {
            continue
        }

        transformerList = transformerList.concat(sub.transformers)
    }

    return transformerList
}

export function getStatusLabel(status) {
    if (!status) {
        return "Active"
    }

    return "Inactive"
}