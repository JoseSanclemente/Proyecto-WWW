import uuid
from django.db import models

# Create your models here.


class Contract(models.Model):
    contract_id = models.UUIDField(primary_key=True, unique=True, default=uuid.uuid4, editable=False)
    client_id = models.ForeignKey('client.Client', null=True, on_delete=models.SET_NULL)
    predial = models.FloatField()
    address = models.CharField(max_length=30)
    meter = models.ForeignKey('meter.Meter', null=True, on_delete=models.SET_NULL)
    active = models.BooleanField()

