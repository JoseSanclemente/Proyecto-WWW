from rest_framework import routers

from .viewsets import ClientViewSets

router = routers.SimpleRouter()
router.register('client', ClientViewSets)

urlpatterns = router.urls
