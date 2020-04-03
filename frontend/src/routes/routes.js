import DashboardLayout from "@/pages/Layout/DashboardLayout.vue";

import Dashboard from "@/pages/Dashboard.vue";
import UserProfile from "@/pages/UserProfile.vue";
import TableList from "@/pages/TableList.vue";
import Typography from "@/pages/Typography.vue";
import Icons from "@/pages/Icons.vue";
import Maps from "@/pages/Maps.vue";
import Notifications from "@/pages/Notifications.vue";
import UserManagement from "@/pages/UserManagement.vue";
import AssetsManagement from "@/pages/AssetsManagement.vue";
import ClientsManagement from "@/pages/ClientsManagement.vue";

const routes = [
  {
    path: "/",
    component: DashboardLayout,
    redirect: "/user-management",
    children: [
      {
        path: "dashboard",
        name: "Dashboard",
        component: Dashboard
      },
      {
        path: "user-management",
        name: "User Management",
        component: UserManagement
      },
      {
        path: "clients-management",
        name: "Clients Management",
        component: ClientsManagement
      },
      {
        path: "assets-management",
        name: "Assets Management",
        component: AssetsManagement
      },
      {
        path: "user",
        name: "User Profile",
        component: UserProfile
      },
      {
        path: "table",
        name: "Table List",
        component: TableList
      },
      {
        path: "typography",
        name: "Typography",
        component: Typography
      },
      {
        path: "icons",
        name: "Icons",
        component: Icons
      },
      {
        path: "maps",
        name: "Maps",
        meta: {
          hideFooter: true
        },
        component: Maps
      },
      {
        path: "notifications",
        name: "Notifications",
        component: Notifications
      }
    ]
  }
];

export default routes;
