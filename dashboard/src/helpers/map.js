export function markersWithPivots(substations) {
    let polyLines = [];

    for (let i = 0; i < substations.length; i++) {
        let sub = substations[i];
        if (sub.transformers == null) {
            continue
        }

        let pivot = { lat: parseFloat(sub.latitude), lng: parseFloat(sub.longitude) };

        for (let j = 0; j < sub.transformers.length; j++) {
            let trans = sub.transformers[j];
            let pivotChild = { lat: parseFloat(trans.latitude), lng: parseFloat(trans.longitude) };

            let lyne = [
                pivot, pivotChild
            ];
            polyLines.push(lyne);
        }
    }

    return polyLines;
}
