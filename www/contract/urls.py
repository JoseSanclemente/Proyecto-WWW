from rest_framework import routers

from .viewsets import ContractViewSets

router = routers.SimpleRouter()
router.register('contract', ContractViewSets)

urlpatterns = router.urls
