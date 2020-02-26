from rest_framework import routers

from .viewsets import ElectricaribeViewSets

router = routers.SimpleRouter()
router.register('dummy', ElectricaribeViewSets)

urlpatterns = router.urls
