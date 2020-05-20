import api from "./api";

const create = payload => {
  return api.post("consumer", payload);
};

const list = () => {
  return api.get("consumer/list");
};

const createContract = payload => {
  return api.post("consumer/contract", payload)
}

const getConsumerPDF = payload => {
  let url = "consumer/contract/pdf?contract_id=" + payload
  return api.get(url)
}

const listConsumerContracts = payload => {
  let url = "consumer/contract/list?consumer_id=" + payload
  return api.get(url)
}

const payConsumerBill = payload => {
  return api.put("bill/payment/operator", payload)
}

export default { list, create, createContract, getConsumerPDF, listConsumerContracts, payConsumerBill };
