# Generated by Django 3.0.3 on 2020-04-02 18:44

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ('meter', '0002_auto_20200402_1831'),
        ('client', '0001_initial'),
        ('contract', '0001_initial'),
    ]

    operations = [
        migrations.AlterField(
            model_name='contract',
            name='client_id',
            field=models.ForeignKey(null=True, on_delete=django.db.models.deletion.SET_NULL, to='client.Client'),
        ),
        migrations.AlterField(
            model_name='contract',
            name='meter',
            field=models.ForeignKey(null=True, on_delete=django.db.models.deletion.SET_NULL, to='meter.Meter'),
        ),
    ]
