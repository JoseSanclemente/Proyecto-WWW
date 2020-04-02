from rest_framework import routers

from .viewsets import ReadsViewSets

router = routers.SimpleRouter()
router.register('Reads', ReadsViewSets)

urlpatterns = router.urls
