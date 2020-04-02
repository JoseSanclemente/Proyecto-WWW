from django.db import models
import uuid

# Create your models here.
class Meter(models.Model):
    meter_id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    transformer = models.ForeignKey('transformer.Transformer', null=True, on_delete=models.SET_NULL)
    metmodel = models.ForeignKey('metmodel.Metmodel', null=True, on_delete=models.SET_NULL)
    address = models.CharField(max_length=30)
    active = models.BooleanField()
