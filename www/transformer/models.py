from django.db import models

# Create your models here.
class Transformer(models.Model):
    trans_id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    subestation = models.ForeignKey('subestation.Subestation', null=True, on_delete=models.SET_NULL)
    name = models.CharField(max_length=30)
    longitude = models.DecimalField(max_digits=30, decimal_places=15)
    latitude = models.DecimalField(max_digits=30, decimal_places=15)