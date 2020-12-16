package team01.yowayowa.sekibutton

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import android.util.Log;
import android.view.MenuItem

class SettingsActivity:AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_settings)
        supportActionBar?.setDisplayHomeAsUpEnabled(true)//戻るボタンの有効化

        // 設定用の Fragment を表示
        supportFragmentManager
            .beginTransaction()
            .replace(R.id.settingsContainer, PreferenceFragment())
            .commit()
    }

    override fun onOptionsItemSelected(item: MenuItem): Boolean {
        finish()
        return super.onOptionsItemSelected(item)
    }
}