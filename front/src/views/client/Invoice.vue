<template>
  <div>
    <div class="form-check-out">
      <router-link to="/history" class="back">Regresar al historial</router-link>
      <div class="form-background">
        <h3 class="title">Factura</h3>
        <div class="modal-body">
          <table class="history-table" align="left" v-if="loaded">
            <tr>
              <td align="right">NIT:</td>
              <td align="left">830047819</td>
            </tr>
            <tr>
              <td align="right">Tel:</td>
              <td align="left">677 88 99</td>
            </tr>
            <tr>
              <td align="right">Fecha:</td>
              <td align="left">{{ invoice.info.date.split(" ")[0] }}</td>
            </tr>
            <tr>
              <td align="right">Entidad aprobadora:</td>
              <td align="left">{{invoice.info.approval_entity}}</td>
            </tr>
            <tr>
              <td align="right">Número de aprobación:</td>
              <td align="left">{{invoice.info.approval_number}}</td>
            </tr>
            <tr>
              <td align="right">Descuento:</td>
              <td align="left">${{invoice.info.discount}}</td>
            </tr>
            <tr>
              <td align="right">Total:</td>
              <td align="left">${{invoice.info.total}}</td>
            </tr>
            <tr>
              <td align="right">Total IVA:</td>
              <td align="left">${{invoice.info.total_iva}}</td>
            </tr>
            <tr>
              <td align="right">Total pagado:</td>
              <td align="left">${{invoice.info.total_payed}}</td>
            </tr>
            <tr>
              <td colspan="3">
                <table class="history-table" align="left">
                  <tr>
                    <th>ID del producto</th>
                    <th>Precio</th>
                    <th>Cantidad</th>
                    <th>IVA</th>
                  </tr>
                  <tr v-for="product in invoice.products" :key="product.product_id">
                    <td>{{ product.product_id }}</td>
                    <td>${{ product.price }}</td>
                    <td>{{ product.quantity }}</td>
                    <td>{{ product.iva }}%</td>
                  </tr>
                </table>
              </td>
            </tr>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { notify } from '@/notify.js'
import { getInvoice } from '@/api/api.js'

export default {
  data: function () {
    return {
      invoice: {},
      loaded: false
    }
  },

  mounted () {
    getInvoice(this.$route.params.invoice_id)
      .then(response => {
        this.invoice = response.data
        console.log(this.invoice)
        this.loaded = true
      })
      .catch(() => {
        notify({
          body: 'No se pudo cargar la factura',
          type: 'error'
        })
      })
  }
}
</script>

<style scoped>
.form-check-out {
  display: flex;
  align-items: center;
  flex-direction: column;
  height: 100vh;
  font-size: 18px;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}

.title {
  color: white;
  margin: 10px;
}

.modal-body {
  margin: 20px 0;
  color: white;
  max-height: 200px;
  display: block;
}

.history-table {
  border-collapse: collapse;
  display: block;
}

.history-table th {
  padding: 10px;
}

.history-table td {
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
