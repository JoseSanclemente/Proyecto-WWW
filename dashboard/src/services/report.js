import api from "./api";

const topConsumers = payload => {
    return api.post("report/consumer/top", payload);
};

const annualConsumption = () => {
    return api.post("report/yearly");
};

const monthlyConsumption = payload => {
    return api.post("report/monthly", payload)
}

export default { topConsumers, annualConsumption, monthlyConsumption };
