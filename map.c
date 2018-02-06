#include <gtk/gtk.h>
#include <osm-gps-map.h>

int main (int argc, char **argv)
{
	gtk_init(&argc, &argv);
	
	GtkWidget *map = g_object_new (OSM_TYPE_GPS_MAP,
                        //"map-source",opt_map_provider,
                        //"tile-cache",cachedir,
                        //"tile-cache-base", cachebasedir,
                        "proxy-uri",g_getenv("http_proxy"),
                        NULL);

	GtkWidget *osd = g_object_new (OSM_TYPE_GPS_MAP_OSD,
                        "show-scale",TRUE,
                        "show-coordinates",TRUE,
                        "show-crosshair",TRUE,
                        "show-dpad",TRUE,
                        "show-zoom",TRUE,
                        "show-gps-in-dpad",TRUE,
                        "show-gps-in-zoom",TRUE,
                        "dpad-radius", 30,
                        NULL);
	osm_gps_map_layer_add(OSM_GPS_MAP(map), OSM_GPS_MAP_LAYER(osd));
	GtkWidget *window = gtk_window_new (GTK_WINDOW_TOPLEVEL);
	g_signal_connect (window, "destroy", G_CALLBACK(gtk_main_quit), NULL);
	GtkWidget *layout = gtk_box_new(GTK_ORIENTATION_VERTICAL, 10);
	gtk_container_add(GTK_CONTAINER(window), GTK_WIDGET(layout));
	gtk_box_pack_start(GTK_BOX(layout), map, TRUE, TRUE, 0);
	gtk_widget_show_all(window);
	gtk_main();

	return 0;
}
