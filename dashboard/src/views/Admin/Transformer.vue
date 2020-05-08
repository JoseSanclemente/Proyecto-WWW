<template>
  <div>
    <section>
      <transformer-form></transformer-form>
    </section>
    <section></section>
  </div>
</template>

<script>
import TransformerForm from "@/components/transformer/TransformerForm.vue";

export default {
  name: "Transformer",
  components: {
    TransformerForm
  }
};
</script>

<style scoped>
section {
  margin: 1.2em 0;
}
</style>