import uuid
from django.db import models

# Create your models here.


class Invoice(models.Model):
    invoice_id = models.UUIDField(primary_key=True, unique=True, default=uuid.uuid4, editable=False)
    contract_id = models.UUIDField(editable=False)
    period_id = models.UUIDField(editable=False)

