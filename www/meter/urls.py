from rest_framework import routers

from .viewsets import MeterViewSets

router = routers.SimpleRouter()
router.register('Meter', MeterViewSets)

urlpatterns = router.urls
