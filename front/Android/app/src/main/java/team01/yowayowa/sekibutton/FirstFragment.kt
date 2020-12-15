package team01.yowayowa.sekibutton

import android.Manifest
import android.content.Context
import android.content.pm.PackageManager
import android.hardware.camera2.CameraAccessException
import android.hardware.camera2.CameraManager
import android.hardware.camera2.CameraManager.TorchCallback
import android.os.Bundle
import android.os.Handler
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ImageButton
import android.widget.Toast
import androidx.core.content.ContextCompat.checkSelfPermission
import androidx.fragment.app.Fragment
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch


/**
 * A simple [Fragment] subclass as the default destination in the navigation.
 */
class FirstFragment : Fragment() {

    private lateinit var  McameraManager : CameraManager
    private var McameraID: String? = null
    private var SW : Boolean = false

    override fun onCreateView(
            inflater: LayoutInflater, container: ViewGroup?,
            savedInstanceState: Bundle?
    ): View? {
        // Inflate the layout for this fragment
        return inflater.inflate(R.layout.fragment_first, container, false)
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        McameraManager = activity?.getSystemService(Context.CAMERA_SERVICE) as CameraManager
        //ライト点灯用
        McameraManager.registerTorchCallback(object : TorchCallback() {
            override fun onTorchModeChanged(
                cameraId: String,
                enabled: Boolean
            ) {
                super.onTorchModeChanged(cameraId, enabled)
                McameraID = cameraId
                SW = enabled
            }
        }, Handler())

        view.findViewById<ImageButton>(R.id.imageButton).setOnClickListener {
            requestCameraPermission()
        }
    }

    private fun requestCameraPermission() {
        // カメラ権限あり
        if (checkSelfPermission(requireContext(),Manifest.permission.CAMERA) == PackageManager.PERMISSION_GRANTED) {
            flushTorch()
            return
        }
        // カメラ権限なし
        if (shouldShowRequestPermissionRationale(Manifest.permission.CAMERA)) {
            Toast.makeText(context,"権限なし",Toast.LENGTH_SHORT).show()
            return
        }
        requestPermissions(arrayOf(Manifest.permission.CAMERA), 200)
    }

    private fun flushTorch(){
        if (McameraID == null) {
            return
        }
        try {
            //非同期でカメラを5回点滅させる
            GlobalScope.launch{
                for (i in 1..10){
                    if (!SW) {
                        McameraManager.setTorchMode(McameraID!!, true)
                    } else {
                        McameraManager.setTorchMode(McameraID!!, false)
                    }
                    delay(100L)
                }
            }
        } catch (e: CameraAccessException) {
            //エラー処理
            e.printStackTrace()
        }
    }
}