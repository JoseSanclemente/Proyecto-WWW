import uuid
from django.db import models

# Create your models here.


class Invoice(models.Model):
    invoice_id = models.UUIDField(primary_key=True, unique=True, default=uuid.uuid4, editable=False)
    contract_id = models.ForeignKey('contract.Contract', null=True, on_delete=models.SET_NULL)
    period_id = models.ForeignKey('period.Period', null=True, on_delete=models.SET_NULL)

