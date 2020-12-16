package team01.yowayowa.sekibutton

import android.content.Intent
import android.os.Bundle
import android.provider.Settings
import com.google.android.material.floatingactionbutton.FloatingActionButton
import com.google.android.material.snackbar.Snackbar
import androidx.appcompat.app.AppCompatActivity
import android.view.Menu
import android.view.MenuItem
import androidx.viewpager2.widget.ViewPager2
import com.google.android.material.tabs.TabLayout
import com.google.android.material.tabs.TabLayoutMediator

class MainActivity : AppCompatActivity() {

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
}