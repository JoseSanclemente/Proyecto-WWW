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