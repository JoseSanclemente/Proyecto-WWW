from rest_framework import routers

from .viewsets import MetmodelViewSets

router = routers.SimpleRouter()
router.register('Metmodel', MetmodelViewSets)

urlpatterns = router.urls
