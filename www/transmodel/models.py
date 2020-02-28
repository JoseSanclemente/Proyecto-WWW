from django.db import models
import uuid

# Create your models here.
class Transmodel(models.Model):
    model_id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    name = models.CharField(max_length=30)
    manufacturer = models.CharField(max_length=30)
    serie = manufacturer = models.CharField(max_length=15)
    primary_voltage = models.FloatField()
    secundary_voltage = models.FloatField()
    stream_output = models.FloatField()
    weight = models.FloatField()
    height = models.FloatField()
    width = models.FloatField()
    depth = models.FloatField()
