from django.db import models
import uuid

# Create your models here.
class Metmodel(models.Model):
    metmodel_id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    name = models.CharField(max_length=30)
    weight = models.FloatField()
    precision = models.FloatField()
    frequency = models.FloatField()
    power_factor = models.FloatField()
    rated_current = models.FloatField()
