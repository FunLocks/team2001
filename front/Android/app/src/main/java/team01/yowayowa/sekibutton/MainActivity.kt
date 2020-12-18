package team01.yowayowa.sekibutton

import android.Manifest
import android.content.Context
import android.content.Intent
import android.content.pm.PackageManager
import android.hardware.camera2.CameraManager
import android.location.Location
import android.location.LocationListener
import android.location.LocationManager
import android.os.Bundle
import android.provider.Settings
import com.google.android.material.floatingactionbutton.FloatingActionButton
import com.google.android.material.snackbar.Snackbar
import androidx.appcompat.app.AppCompatActivity
import android.view.Menu
import android.view.MenuItem
import android.widget.Toast
import androidx.core.app.ActivityCompat
import androidx.viewpager2.widget.ViewPager2
import com.google.android.material.tabs.TabLayout
import com.google.android.material.tabs.TabLayoutMediator
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaTypeOrNull
import okhttp3.RequestBody.Companion.toRequestBody
import org.json.JSONObject
import java.io.IOException
import java.util.*

class MainActivity : AppCompatActivity(),LocationListener {

    lateinit var mLocationManager : LocationManager
    private var myLocate : Location? = null

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        setSupportActionBar(findViewById(R.id.toolbar))

        //viewPager & indicator
        val viewPager = findViewById<ViewPager2>(R.id.viewPager2)
        viewPager.adapter = PagerAdapter(this)
        viewPager.orientation = ViewPager2.ORIENTATION_HORIZONTAL
        val indicator = findViewById<TabLayout>(R.id.indicator)
        TabLayoutMediator(indicator, viewPager) { _, _ -> }.attach()

        checkLocationPermission()
        //初回のみ
        if (PreferencesUtil().isFirstJudgment(this)){
            PreferencesUtil().setFirstFlag(this)
            val intent = Intent(this, WalkThroughActivity::class.java)
            //帰ってこれなくする
            intent.flags = Intent.FLAG_ACTIVITY_CLEAR_TASK or Intent.FLAG_ACTIVITY_NEW_TASK
            startActivity(intent)
        }
    }

    override fun onCreateOptionsMenu(menu: Menu): Boolean {
        // Inflate the menu; this adds items to the action bar if it is present.
        menuInflater.inflate(R.menu.menu_main, menu)
        return true
    }

    override fun onOptionsItemSelected(item: MenuItem): Boolean {
        // Handle action bar item clicks here. The action bar will
        // automatically handle clicks on the Home/Up button, so long
        // as you specify a parent activity in AndroidManifest.xml.
        return when (item.itemId) {
            R.id.action_settings -> {
                openPreferenceActivity()
                true
            }
            else -> super.onOptionsItemSelected(item)
        }
    }

    private fun openPreferenceActivity(){
        val intent = Intent(applicationContext,SettingsActivity::class.java)
        startActivity(intent)
    }

    private val permissionsRequestCode:Int = 1000;
    //権限周り
    private fun checkLocationPermission() {
        if (ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_FINE_LOCATION)
            != PackageManager.PERMISSION_GRANTED && ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_COARSE_LOCATION)
            != PackageManager.PERMISSION_GRANTED) { // パーミッションの許可を取得する

            ActivityCompat.requestPermissions(this, arrayOf(
                Manifest.permission.ACCESS_FINE_LOCATION,
                Manifest.permission.ACCESS_COARSE_LOCATION
            ),
                permissionsRequestCode
            )
        }else locationStart()
    }

    //パーミッション許可を乞うダイアログから与えられた応答に対するリアクション
    override fun onRequestPermissionsResult(requestCode: Int,
                                            permissions: Array<String>, grantResults: IntArray) {
        when (requestCode) {
            permissionsRequestCode -> {
                // If request is cancelled, the result arrays are empty.
                if ((grantResults.isNotEmpty() && grantResults[0] == PackageManager.PERMISSION_GRANTED)) {
                    locationStart()
                    Toast.makeText(applicationContext,"gpsの使用許可が下りました。位置情報を取得します。",Toast.LENGTH_SHORT).show()
                    Toast.makeText(applicationContext,myLocate?.latitude.toString(),Toast.LENGTH_SHORT).show()
                } else {
                    Toast.makeText(applicationContext,"gpsの使用許可が下りませんでした。位置情報を取得できません。",Toast.LENGTH_SHORT).show()
                }
                return
            }
            else ->{

            }
        }
    }

    //位置情報を取得
    private fun locationStart(){
        mLocationManager =
            getSystemService(Context.LOCATION_SERVICE) as LocationManager
        if(ActivityCompat.checkSelfPermission(this,Manifest.permission.ACCESS_FINE_LOCATION) == 0){
            when {
                mLocationManager.isProviderEnabled(LocationManager.GPS_PROVIDER) -> {
                    myLocate = mLocationManager!!.getLastKnownLocation("gps")
                    mLocationManager.requestLocationUpdates("gps",1000,10F,this)
                }
                mLocationManager.isProviderEnabled(LocationManager.NETWORK_PROVIDER) -> {
                    myLocate = mLocationManager!!.getLastKnownLocation(LocationManager.NETWORK_PROVIDER)
                    mLocationManager.requestLocationUpdates(LocationManager.NETWORK_PROVIDER,1000,10F,this)
                }
                else -> {
                    //GPSが取れなかった時の処理
                    return
                }
            }
        }
    }


    internal fun sendToServer(){
        GlobalScope.launch {
            val url: String = "http://153.120.166.49:8080/ahchoo/post"
            val client: OkHttpClient = OkHttpClient.Builder().build()

            // create json
            val json = JSONObject()
            json.put("latitude", myLocate?.latitude.toString())
            json.put("longitude", myLocate?.longitude.toString())

            // post
            val postBody =
                json.toString().toRequestBody("application/json; charset=utf-8".toMediaTypeOrNull())
            val request: Request = Request.Builder().url(url).post(postBody).build()
            val response = client.newCall(request).execute()

            val result: String? = response.body?.string()
            response.close()
        }
    }

    override fun onLocationChanged(location: Location?) {
        myLocate = location
    }

    override fun onStatusChanged(provider: String?, status: Int, extras: Bundle?) {
        //TODO("Not yet implemented")
    }

    override fun onProviderEnabled(provider: String?) {
        //TODO("Not yet implemented")
    }

    override fun onProviderDisabled(provider: String?) {
        //TODO("Not yet implemented")
    }
}