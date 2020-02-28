import uuid
from django.db import models

# Create your models here.


class Contract(models.Model):
    contract_id = models.UUIDField(primary_key=True, unique=True, default=uuid.uuid4, editable=False)
    client_id = models.UUIDField(editable=False)
    predial = models.FloatField()
    address = models.CharField(max_length=30)
    meter = models.UUIDField()
    active = models.BooleanField()

