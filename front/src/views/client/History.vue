<template>
  <div>
    <div class="form-check-out">
      <router-link to="/user" class="back">Regresar al menú principal</router-link>
      <div class="form-background">
        <div class="form">
          <h3 class="title">Historial de compras</h3>
          <div class="modal-body">
            <table class="history-table" align="center">
              <tr>
                <th>Fecha</th>
                <th>ID de factura</th>
                <th>Total</th>
                <th>IVA</th>
                <th>Total pagado</th>
                <th>Open</th>
              </tr>
              <tr v-for="invoice in history" :key="invoice.invoice_id">
                <td>{{ invoice.approval_date.split("T")[0] }}</td>
                <td>{{ invoice.invoice_id }}</td>
                <td>${{ invoice.total }}</td>
                <td>${{ invoice.total_iva }}</td>
                <td>${{ invoice.total_payed }}</td>
                <td>
                  <div class="confirmation-button" @click="goToInvoice(invoice.invoice_id)">Open</div>
                </td>
              </tr>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { notify } from '@/notify.js'
import { getHistory } from '@/api/api.js'

export default {
  data: function () {
    return {
      history: {}
    }
  },

  methods: {
    goToInvoice (invoiceID) {
      this.$router.push({
        name: 'invoice',
        params: {
          invoice_id: invoiceID
        }
      })
    }
  },

  mounted () {
    getHistory()
      .then(response => {
        this.history = response.data
      })
      .catch(() => {
        notify({
          body: 'No se pudo cargar el historial de facturas',
          type: 'error'
        })
      })
  }
}
</script>

<style scoped>
.form-check-out {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 125vh;
  font-size: 18px;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}

.title {
  color: white;
  margin: 10px;
}

.product-image {
  max-width: 400px;
  max-height: 400px;

  float: left;
}

.product-information {
  color: white;
  display: block;
  width: 400px;
}

.product-table {
  margin: 20px;
}

.modal-container {
  width: 300px;
  margin: 0px auto;
  padding: 20px 30px;
  background-color: black;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
  font-family: Helvetica, Arial, sans-serif;
  color: white;
}

.modal-header h3 {
  margin-top: 0;
  color: #42b983;
}

.modal-body {
  margin: 20px 0;
  color: white;
  max-height: 200px;
  overflow: auto;
}

.history-table {
  border: 2px solid white;
  border-radius: 10px;
  border-collapse: collapse;
}

.history-table th {
  border: 2px solid white;
  padding: 10px;
}

.history-table td {
  border: 2px solid white;
  padding: 10px;
}

.item-image {
  max-width: 68px;
  max-height: 68px;
}

.controls-container {
  width: 30px;
  margin: auto;
}
</style>
