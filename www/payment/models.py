from django.db import models
import uuid

# Create your models here.
class Payment(models.Model):
    payment_id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    invoice_id = models.ForeignKey('invoice.Invoice', null=True, on_delete=models.SET_NULL)
    paid_amount = models.FloatField()